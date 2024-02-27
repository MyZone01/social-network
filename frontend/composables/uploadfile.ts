
export const sendFile = async (file:File, token: string) => {
    const body = new FormData();
    body.append('file', file, file.name);

    // Think of Using decoder on the token from /utils/transformer.js
    return await fetch('/api/upload', {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body,
    });
}