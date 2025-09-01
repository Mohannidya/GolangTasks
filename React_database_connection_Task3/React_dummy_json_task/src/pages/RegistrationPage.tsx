import { useEffect, useState } from 'react'


interface User {
  name: string
  email: string
  password: string
}

function RegistraionPage() {
  const [form, setForm] = useState<User>({
    name: '',
    email: '',
    password: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const res = await fetch('http://localhost:8080/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form),
    });
    const data = await res.json();
    console.log(data);
    alert("User added!");
  };

  return (
        <section className="max-w-md mx-auto">
      <h1 className="text-2xl font-bold mb-4">Registration</h1>
    <div className="app">
      <h2>Create User</h2>
      <form onSubmit={handleSubmit}>
        <div>
         <label className="block text-sm font-medium mb-1" htmlFor="name">Name</label>
  <input
  type="text"
  name="name"
   className="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Your full name"
  value={form.name}
  onChange={handleChange}
  required
/>
</div>

<div>
  <label className="block text-sm font-medium mb-1" htmlFor="email">Email</label>
<input
  type="email"
  name="email"
 className="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="you@example.com"
  value={form.email}
  onChange={handleChange}
  
  required
/>
</div>

<div>
   <label className="block text-sm font-medium mb-1" htmlFor="password">Password</label>
<input
  type="password"
  name="password"
   className="w-full rounded-md border px-3 py-2 outline-none focus:ring-2 focus:ring-blue-500"
  placeholder="Password"
  value={form.password}
  onChange={handleChange}
  required
/>

</div>
        <button 
        type="submit" 
        className="w-full rounded-md bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700 disabled:opacity-60">Add User</button>
      </form>
    </div>
    </section>
  );
}

export default RegistraionPage;
