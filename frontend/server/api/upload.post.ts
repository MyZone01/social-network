

export default defineEventHandler(async (event) => {
    console.log("Uploading file");

    const body = await readMultipartFormData(event);
    console.log(body);
    return {
        status: 200,
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: 'File uploaded' }),
    };
});