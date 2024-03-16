import type { Event } from "~/types";

export const useEvents = () => {
    async function createEvent(eventDetail: any, groupId: string) {
        const { data } = await useFetch("/api/groups/events/create", {
            method: "post",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,
            body: JSON.stringify(eventDetail),
            query: {
                gid: groupId,
            }
        })
        return JSON.parse(data.value as string)
    }

    const getAllEvents = async (groupId: string) => {
        const response = await $fetch<any>("/api/groups/events", {
            headers: useRequestHeaders(["cookie"]) as HeadersInit,
            query: {
                q: 1,
                u: 1,
                gid: groupId
            }
        });

        return response.data;
    };

    async function joinEvent(eventid: string) {
        const { data, error } = await useFetch("/api/groups/request/join", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,
            query: {
                eid: eventid,
            },
        });

        return { data, error };
    }

    return { createEvent, getAllEvents }
}