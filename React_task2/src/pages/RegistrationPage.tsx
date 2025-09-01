// src/pages/RegistrationPage.tsx
import { useState } from "react";
import { useNavigate } from "react-router-dom";

type FormData = {
  name: string;
  email: string;
  password: string;
};

type CreatedUser = {
  id: string | number;
  name: string;
  email: string;
  // ...any other fields your backend returns
};

export default function RegistrationPage() {
  const [form, setForm] = useState<FormData>({ name: "", email: "", password: "" });
  const [error, setError] = useState<string | null>(null);
  const [ loading, setLoading] = useState(false);
  const navigate = useNavigate();

  function onChange(e: React.ChangeEvent<HTMLInputElement>) {
    const { name, value } = e.target;
    setForm((f) => ({ ...f, [name]: value }));
  }

  async function onSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setError(null);

    // basic client-side checks
    if (!form.name.trim()) return setError("Name is required");
    if (!/^\S+@\S+\.\S+$/.test(form.email)) return setError("Enter a valid email");
    if (form.password.length < 6) return setError("Password must be ≥ 6 characters");

    try {
      setLoading(true);
      const res = await fetch("http://localhost:58080/users", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(form),
      });

      if (!res.ok) {
        // Try to read an error message from the server
        let msg = `Request failed: ${res.status} ${res.statusText}`;
        try {
          const data = await res.json();
          if (data?.error) msg = String(data.error);
          if (data?.message) msg = String(data.message);
        } catch { /* ignore JSON parse errors */ }
        throw new Error(msg);
      }

      const created: CreatedUser = await res.json();
      console.log("Created user:", created);

      // go to users list (adjust if you use a different route)
      navigate("/users", { replace: true });
    } catch (err: any) {
      setError(err?.message ?? "Something went wrong");
    } finally {
      setLoading(false);
    }
  }

  return (
    <section className="max-w-md mx-auto">
      <h1 className="text-2xl font-bold mb-4">Registration</h1>

      {error && (
        <div className="mb-3 rounded-md border border-red-300 bg-red-50 px-3 py-2 text-sm text-red-700">
          {error}
        </div>
      )}

      <form onSubmit={onSubmit} className="space-y-3">
        <div>
          <label className="block text-sm font-medium mb-1" htmlFor="name">Name</label>
          <input
            id="name"
            name="name"
            value={form.name}
            onChange={onChange}
              className="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Your full name"
            required
          />
        </div>

        <div>
          <label className="block text-sm font-medium mb-1" htmlFor="email">Email</label>
          <input
            id="email"
            name="email"
            type="email"
            value={form.email}
            onChange={onChange}
            className="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="you@example.com"
            required
          />
        </div>

        <div>
          <label className="block text-sm font-medium mb-1" htmlFor="password">Password</label>
          <input
            id="password"
            name="password"
            type="password"
            value={form.password}
            onChange={onChange}
            className="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="••••••••"
            required
            minLength={6}
          />
        </div>

        <button
          type="submit"
          disabled={loading}
          className="w-full rounded-md bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700 disabled:opacity-60"
        >
          {loading ? "Submitting..." : "Submit"}
        </button>
      </form>
    </section>
  );
}
