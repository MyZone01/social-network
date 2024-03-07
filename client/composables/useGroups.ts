import type { Group } from "~/types";

export const useGroups = () => {
  const getAllGroups = async () => {

    const response = await $fetch<Group[]>("/api/groups", {
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
    });

    return response;
  };

  const getGroupByID = async (id: string) => {
    const response = await $fetch<Group>("/api/groups/" + id, {
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
    });

    return response;
  };

  return {
    getAllGroups,
    getGroupByID
  };
}
