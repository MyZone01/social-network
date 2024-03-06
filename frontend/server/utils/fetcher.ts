export const fetcher = async (url: string, method: string, body: any, token: string) => {
    try {
        const response = await fetch(url, {
            method: method,
            body,
            headers: {
                // Think of Using decoder on the token from /utils/transformer.js
                Authorization: `Bearer ${token}`,
            },
        })
        return response.json()
    } catch (e) {
        return {
            status: 500,
            body: 'Internal server error',
        }
    }
}