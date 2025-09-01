import type { User } from "../pages/UsersPage"

export default function SingleUser(user: User){
return (
    <>
     <div
            key={user.id}
            className="p-4 border rounded-lg shadow hover:shadow-lg transition"
          >
            <img
              src={user.image}
              alt={`${user.firstName} ${user.lastName}`}
              className="w-20 h-20 rounded-full mx-auto"
            />
            <h3 className="mt-2 text-lg font-semibold text-center">
              {user.firstName} {user.lastName}
            </h3>
            <p className="text-gray-600 text-center">{user.email}</p>
          </div>
    </>
)
}