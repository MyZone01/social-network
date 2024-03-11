import { GroupMember } from '~/types'

export const useGroupRequest = () => {
    async function joinRequest(groupId: string | undefined): Promise<{ data: any; error: any; }> {

        const { data, error } = await useFetch("/api/group/request/join", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gid: groupId,
            },
        });
        return { data, error };
    }

    async function getJoinRequests(
        groupId: string | undefined
    ): Promise<GroupMember[] | null> {

        const data = await $fetch("/api/group/request/join-requests", {
            method: "GET",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gid: groupId,
            },
        });
        return data;
    }

    async function acceptJoinRequest(
        gId: string,
        rId: string
    ): Promise<any> {

        const data = await $fetch("/api/group/request/accept", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gId,
                rId,
            },
        });
        return { data };
    }

    async function declneJoinRequest(
        gId: string,
        rId: string
    ): Promise<any> {

        const data = await $fetch("/api/group/request/decline", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gId,
                rId,
            },
        });
        return { data };
    }

    async function getUserGroups() {

        const response = await $fetch("/api/user/groups", {
            method: "GET",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

        });

        return response
    }

    return {
        joinRequest,
        getUserGroups,
        declneJoinRequest,
        acceptJoinRequest,
        getJoinRequests
    }

}