import NavBar from './components/NavBar.tsx'
import HomePage from './pages/HomePage.tsx'
import RegistrationPage from './pages/RegistrationPage.tsx'
import UsersPage from './pages/UsersPage.tsx'
import { Route, Routes } from "react-router-dom";

export default function Myapp() {
  return (
    <>
      <NavBar />
      <div className="mx-auto max-w-5xl px-4 py-6">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/users" element={<UsersPage />} />
          <Route path="/registration" element={<RegistrationPage />} />
          <Route path="*" element={<h2 className="p-6">404 — Not Found</h2>} />
        </Routes>
      </div>
    </>
  );
}