import HolidayCalendar from "./components/HolidayCalendar";

export const BASE_URL=import.meta.env.MODE = "development" ? "http://localhost:8080/holidays" : "/api";

function App() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-indigo-500 to-black text-white">
      <h1 className="text-5xl font-extrabold mb-8 drop-shadow-lg">📅Calendar</h1>
      <HolidayCalendar />
    </div>
  );
}

export default App;
