import { getUserFromToken } from "~/server/utils/session";

export default defineEventHandler(async (event) => {
  const user = await getUserFromToken(event);
  if (user) event.context.user = user;
});