export default (url, options = {}) => {
    const store = useGlobalAuthStore();

    return $fetch(url, {
        ...options,
        headers: {
            ...options.headers,
            Authorization: `Bearer ${store.token}`
        }
    })
}