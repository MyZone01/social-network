export default defineEventHandler(async (event) => {
  const group = await $fetch("http://localhost:8081/get-all-groups/");
  return group;
});
