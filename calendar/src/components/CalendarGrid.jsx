import { eachDayOfInterval, startOfMonth, endOfMonth, startOfWeek, endOfWeek, format, isSameMonth } from "date-fns";
import DateCell from "./DateCell";

const CalendarGrid = ({ currentMonth, holidays, fetchHolidays, setHolidays }) => {  
  const startDay = startOfWeek(startOfMonth(currentMonth), { weekStartsOn: 0 });
  const endDay = endOfWeek(endOfMonth(currentMonth), { weekStartsOn: 0 });
  const days = eachDayOfInterval({ start: startDay, end: endDay });

  const weekDays = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];

  return (
    <div className="border rounded-lg overflow-hidden shadow-xl bg-white/40 p-4 backdrop-blur-md">
      {/* Weekdays Header */}
      <div className="grid grid-cols-7 text-center font-bold text-white bg-indigo-700 p-2 rounded-t-lg">
        {weekDays.map((day) => (
          <div key={day} className="py-2">{day}</div>
        ))}
      </div>
      
      <div className="grid grid-cols-7 bg-gray-200">
        {days.map((day) => (
          isSameMonth(day, currentMonth) ? (
            <DateCell 
              key={format(day, "yyyy-MM-dd")} 
              day={day} 
              holidays={holidays} 
              fetchHolidays={fetchHolidays} 
              setHolidays={setHolidays}  
            />
          ) : (
            <div key={format(day, "yyyy-MM-dd")} className="p-4 border text-center bg-gray-300 rounded-lg"></div>
          )
        ))}
      </div>
    </div>
  );
};

export default CalendarGrid;
