import { useState } from "react";
import { format } from "date-fns";

const DateCell = ({ day, holidays, fetchHolidays, setHolidays }) => { 
  const [hovered, setHovered] = useState(false);

  const formattedDate = format(day, "yyyy-MM-dd");
  const holiday = holidays[formattedDate]; 

  const addHoliday = async () => {
    const name = prompt("Enter holiday name:");
    const country = prompt("Enter country name:");
    if (!name || !country) return;

    try {
      console.log("Sending request to backend...");
      // Updated to use the /api prefix.
      const response = await fetch("/api/holidays", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ date: formattedDate, name, country }),
      });

      console.log("Response status:", response.status);
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json();
      console.log("Holiday added successfully:", data);

      fetchHolidays(); 
    } catch (error) {
      console.error("Error adding holiday:", error);
    }
  };

  const removeHoliday = async () => {
    if (!holiday?.id) {
      console.error("Error: No holiday ID found for deletion.");
      return;
    }

    try {
      console.log(`Deleting holiday with ID: ${holiday.id}`);
      // Updated to use the /api prefix.
      const response = await fetch(`/api/holidays/${holiday.id}`, {
        method: "DELETE",
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      console.log("Holiday removed successfully.");

      setHolidays(prevHolidays => {
        const updatedHolidays = { ...prevHolidays };
        delete updatedHolidays[formattedDate]; 
        return updatedHolidays;
      });
    } catch (error) {
      console.error("Error deleting holiday:", error);
    }
  };

  return (
    <div
      className="relative p-4 border text-center bg-gray-100 hover:bg-gray-200 transition-all rounded-lg"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      <span className="text-lg font-semibold text-gray-800">{format(day, "d")}</span>

      {/* Show Add button only if no holiday exists */}
      {hovered && !holiday && (
        <button
          onClick={addHoliday}
          className="absolute bottom-2 left-1/2 transform -translate-x-1/2 bg-blue-500 text-white px-3 py-1 text-xs rounded shadow-md hover:bg-blue-600 transition-all"
        >
          + Holiday
        </button>
      )}

      {holiday && (
        <div className="mt-2 text-sm text-red-500 font-medium flex items-center justify-center gap-2">
          {holiday.name}
          <button
            onClick={removeHoliday}
            className="text-xs text-gray-500 hover:text-gray-700"
          >
            ❌
          </button>
        </div>
      )}
    </div>
  );
};

export default DateCell;
