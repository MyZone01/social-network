
export const sendFile = async (file:File, token: string) => {
    // console.log(file, token);
    
    const body = new FormData();
    body.append('file', file, file.name);
    const reponse = await fetch('/api/upload', {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body,
    });
    console.log(reponse);
    return {
        file,
        token
    }
}