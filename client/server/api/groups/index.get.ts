<<<<<<< HEAD
import type { Group, ServerResponse, User } from "~/types";
=======
import type { User } from "~/types";
>>>>>>> group-timeline

export default defineEventHandler(async (event) => {
  if (!event.context.token) {
    return createError({
      statusCode: 401,
      message: "You don't have the rights to access this resource",
    });
  }

  const token = event.context.token;

  const response = await $fetch<ServerResponse<Group[]>>("http://localhost:8081/get-all-groups?isMemberNeeded=true&isUserNeeded=true", {
    method: "GET",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  return response;
});