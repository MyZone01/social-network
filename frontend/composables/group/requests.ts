export async function joinRequest(groupId: string) {
    const store = useGlobalAuthStore()
    const { error } = await useFetch("/api/group/request/join", {
        method: 'POST',
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

export async function getJoinRequests(groupId: string): Promise<any> {
    const store = useGlobalAuthStore()
    const data = await $fetch("/api/group/request/join-requests", {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${store.token}`
        },
        query: {
            gid: groupId
        }
    })
    return { data }
}