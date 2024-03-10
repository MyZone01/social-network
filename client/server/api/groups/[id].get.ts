import { User } from "~/types";

export default defineEventHandler(async (event) => {
  if (!event.context.token) {
    return createError({
      statusCode: 401,
      message: "You don't have the rights to access this resource",
    });
  }

  const token = event.context.token;
  const groupId = getRouterParam(event, "id");
  console.log(">>>>>>>>>>>>>>>>Endpoint hit", groupId);

  const response = await $fetch<User>(`http://localhost:8081/get-group?group_id=${groupId}&isMemberNeeded=true&isUserNeeded=true`, {
    method: "GET",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  return response;
});