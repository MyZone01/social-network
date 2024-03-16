export default defineNitroPlugin(() => {
    runTask('socket:connect')
    runTask('notif:sendNotif')
})
