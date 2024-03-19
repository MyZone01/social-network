import type { GroupMember } from '~/types';

export const useGroupRequest = () => {
    async function joinRequest(groupId: string | undefined): Promise<{data:any,error:any}> {

        const { data, error } = await useFetch("/api/groups/request/join", {
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

        const response = await $fetch("/api/groups/request/join-requests", {
            method: "GET",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gid: groupId,
            },
        });
        const data = JSON.parse(response as string).data        

        return data as GroupMember[] 
    }

    async function acceptJoinRequest(
        gId: string,
        rId: string
    ): Promise<any> {

        const data = await $fetch("/api/groups/request/accept", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gId,
                rId,
            },
        });
        return { data };
    }

    async function declineJoinRequest(
        gId: string,
        rId: string
    ): Promise<any> {

        const data = await $fetch("/api/groups/request/decline", {
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
        declineJoinRequest,
        acceptJoinRequest,
        getJoinRequests
    }

}