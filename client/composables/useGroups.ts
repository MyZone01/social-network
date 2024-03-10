import type { Group } from "~/types";

export const useGroups = () => {
  const groups = ref<Group[]>([]);

  const createGroup = async (group: { title: string, description: string }) => {
    const response = await $fetch<Group>("/api/groups", {
      method: "POST",
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
      body: {
        title: group.title,
        description: group.description
      },
    });

    groups.value = [...groups.value, response];
  };

  const getAllGroups = async () => {
    const response = await $fetch<Group[]>("/api/groups", {
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
    });

    groups.value = response;
  };

  const getGroupByID = async (id: string) => {
    const response = await $fetch<Group>("/api/groups/" + id, {
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
    });

    return response;
  };

  return {
    groups,
    createGroup,
    getAllGroups,
    getGroupByID
  };
}
