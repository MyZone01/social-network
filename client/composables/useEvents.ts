import type { Event, EventParticipant } from "~/types";

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
                participants: 1,
                user: 1,
                gid: groupId
            }
        });

        return response.data;
    };

    async function respondEvent(event: Event, response: string) {
        const { data, error } = await useFetch("/api/groups/events/respond", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,
            query: {
                eid: event.ID,
                gid: event.GroupID
            },
            body: JSON.stringify({ response })
        });

        console.log(data);
        

        return { data: JSON.parse(data.value as unknown as string) as { data: any, message: string }, error };
    }

    return { createEvent, getAllEvents, respondEvent }
}