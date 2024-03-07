export async function joinRequest(groupId: string | undefined) {
  const store = useGlobalAuthStore();
  const { data, error } = await useFetch("/api/group/request/join", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${store.token}`,
    },
    query: {
      gid: groupId,
    },
    onResponseError({}) {
      return { error };
    },
    onRequestError() {
      return { error };
    },
  });
  return { data, error };
}

export async function getJoinRequests(
  groupId: string | undefined
): Promise<any> {
  const store = useGlobalAuthStore();
  const data = await $fetch("/api/group/request/join-requests", {
    method: "GET",
    headers: {
      Authorization: `Bearer ${store.token}`,
    },
    query: {
      gid: groupId,
    },
  });
  return { data };
}

export async function acceptJoinRequest(
  gId: string,
  rId: string
): Promise<any> {
  const store = useGlobalAuthStore();
  const data = await $fetch("/api/group/request/accept", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${store.token}`,
    },
    query: {
      gId,
      rId,
    },
  });
  return { data };
}

export async function declneJoinRequest(
  gId: string,
  rId: string
): Promise<any> {
  const store = useGlobalAuthStore();
  const data = await $fetch("/api/group/request/decline", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${store.token}`,
    },
    query: {
      gId,
      rId,
    },
  });
  return { data };
}

export async function getUserGroups() {
  const store = useGlobalAuthStore();
  const response = await $fetch("/api/user/groups", {
    method:"GET",
    headers: {
      Authorization: `Bearer ${store.token}`,
    },
  });

  return response
}
