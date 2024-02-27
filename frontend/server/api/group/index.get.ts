export default defineEventHandler(async (event) => {
  const token = event.headers.get('Authorization')
  const groups = await $fetch('http://localhost:8081/get-all-groups', {
    headers: {
      'Authorization': `${token}`
    }

  })
  console.log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>from api", groups);
  return { groups }
})
