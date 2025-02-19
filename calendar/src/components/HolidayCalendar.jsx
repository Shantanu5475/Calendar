import { useState, useEffect } from "react";
import { format, addMonths, subMonths, setMonth, setYear } from "date-fns";
import CalendarGrid from "./CalendarGrid";

const HolidayCalendar = () => {
  const [currentMonth, setCurrentMonth] = useState(new Date());
  const [holidays, setHolidays] = useState({});

  const months = [
    "January", "February", "March", "April", "May", "June", 
    "July", "August", "September", "October", "November", "December"
  ];

  const years = Array.from({ length: 6 }, (_, i) => 2024 + i);

  useEffect(() => {
    fetchHolidays();
  }, [currentMonth]);

  const fetchHolidays = async () => {
    try {
      // Updated to use relative URL
      const response = await fetch("/holidays");
      const data = await response.json();
  
      const holidaysMap = {};
      data.forEach(holiday => {
        holidaysMap[holiday.date] = { 
          name: holiday.name, 
          id: holiday.id || holiday._id 
        };
      });
  
      console.log("Fetched holidays:", holidaysMap);
      setHolidays(holidaysMap);
    } catch (error) {
      console.error("Error fetching holidays:", error);
    }
  };

  const handleMonthChange = (e) => {
    const newMonth = parseInt(e.target.value, 10);
    setCurrentMonth(setMonth(currentMonth, newMonth));
  };

  const handleYearChange = (e) => {
    const newYear = parseInt(e.target.value, 10);
    setCurrentMonth(setYear(currentMonth, newYear));
  };

  return (
    <div className="p-6 w-full max-w-4xl mx-auto text-center bg-white/30 backdrop-blur-lg shadow-2xl rounded-lg border border-white/40">
      {/* Header Section */}
      <div className="flex justify-between items-center mb-6 p-4 bg-white shadow-lg rounded-lg border border-gray-300">
        <button 
          onClick={() => setCurrentMonth(subMonths(currentMonth, 1))} 
          className="p-3 bg-gray-200 hover:bg-gray-400 text-gray-800 font-semibold rounded-lg shadow-md transition-all duration-200 transform hover:scale-105"
        >
          ◀ Prev
        </button>

        <div className="text-3xl font-bold text-indigo-700 drop-shadow-md">
          {months[currentMonth.getMonth()]} {currentMonth.getFullYear()}
        </div>

        <button 
          onClick={() => setCurrentMonth(addMonths(currentMonth, 1))} 
          className="p-3 bg-gray-200 hover:bg-gray-400 text-gray-800 font-semibold rounded-lg shadow-md transition-all duration-200 transform hover:scale-105"
        >
          Next ▶
        </button>
      </div>

      {/* Dropdowns for Month & Year */}
      <div className="flex justify-center gap-4 mb-4">
        <select 
          value={currentMonth.getMonth()} 
          onChange={handleMonthChange} 
          className="p-2 border rounded bg-white text-gray-700 shadow-lg hover:bg-gray-100 transition-all"
        >
          {months.map((month, index) => (
            <option key={index} value={index}>{month}</option>
          ))}
        </select>
        
        <select 
          value={currentMonth.getFullYear()} 
          onChange={handleYearChange} 
          className="p-2 border rounded bg-white text-gray-700 shadow-lg hover:bg-gray-100 transition-all"
        >
          {years.map((year) => (
            <option key={year} value={year}>{year}</option>
          ))}
        </select>
      </div>

      <CalendarGrid 
        currentMonth={currentMonth} 
        holidays={holidays} 
        fetchHolidays={fetchHolidays} 
        setHolidays={setHolidays}  
      />
    </div>
  );
};

export default HolidayCalendar;
