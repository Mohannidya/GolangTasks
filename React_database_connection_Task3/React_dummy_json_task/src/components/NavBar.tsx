import "../css/output.css"
import { Link } from "react-router-dom";
function NavBar() {
    return (
        <>
        <nav className="sticky top-0 z-50 border-b border-gray-200 bg-white/80 backdrop-blur 
           dark:border-gray-800 dark:bg-gray-900/80">
                <div className="mx-auto flex h-16 max-w-screen-xl items-center justify-between px-4">
                    <a href="#" className="flex items-center">
                        <img src="/vite.svg" alt="Logo" className="h-8 w-auto"/>
                            <span className="ml-2 text-xl font-semibold text-gray-900 dark:text-white">Spanlet</span>
                    </a>

                    <ul className="hidden md:flex items-center gap-6 hover:text-red-800 text-gray-700 dark:text-gray-200">
                        <Link to="/" className="hover:text-red-800 dark:hover:text-blue-400">Home</Link>
                        <Link to="/registration" className="hover:text-red-800 dark:hover:text-blue-400">Registration</Link>
                        <Link to="/users" className="hover:text-red-800 dark:hover:text-blue-400">Users</Link>
                    </ul>

                    <div className="flex items-center gap-4">
                        <button id="theme-toggle"
                            className="rounded-md p-2 text-gray-600 hover:bg-gray-200 dark:text-gray-300 dark:hover:bg-gray-800">
                            🌙
                        </button>

                        <button className="md:hidden rounded-md p-2 text-gray-600 hover:bg-gray-200 dark:text-gray-300 dark:hover:bg-gray-800">
                            ☰
                        </button>
                    </div>
                </div>
            </nav>
        </>
    )
}

export default NavBar;