

export default defineEventHandler(async (event) => {
    console.log("Uploading file");

    const reader = await readMultipartFormData(event);
    if (!reader) return { status: 400, body: 'Bad request' };
    let file;
    for (let part of reader) {
        if (part.filename) { // assuming that the file part has a filename property
            file = part;
            break;
        }
    }

    if (!file) {
        return {
            status: 400,
            body: 'No file uploaded',
        };
    }

    // Create a new FormData instance
    const body = new FormData();
    // Append the file to the FormData instance
    body.append('file', new Blob([file.data]), file.filename);

    const token = event.headers.get('Authorization');
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    const response = await fetch('http://localhost:8081/upload', {
        method: 'POST',
        headers: {
            "content-type": "multipart/form-data",
            Authorization: token,
        },
        body,
    });
    console.log("File uploaded", response.status);
    
});