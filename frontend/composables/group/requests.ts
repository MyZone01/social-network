export async function joinRequest(groupId: string) {
    const store = useGlobalAuthStore()
    const { error } = await useFetch("/api/group/request/join", {
        method:'POST',
        headers: {
            Authorization: `Bearer ${store.token}`
        },
        query: {
            gid: groupId
        },
        onResponseError({ }) {
            return error
        },
        onRequestError() {
            return error
        }
    })
}