
export default defineEventHandler(async (event) => {
  const token = getHeader(event, 'Authorization')
  const groups = await $fetch('http://localhost:8081/get-all-groups', {
    headers: {
      Authorization: `${token}`
    }

  })
  return { groups }
})
