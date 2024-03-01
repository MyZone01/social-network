const GetAllFollower = async () => {
    console.log("geting all follower");
    const authStore = useGlobalAuthStore()
    let response = await fetch("http://localhost:8081/getAllFollowers", {
        method: "POST",
        headers: {
            "Authorization": "Bearer " + authStore.token
        }
    })
    let responseInJsonFormat = await response.json().catch(err => ({ error: err }))
    console.log(responseInJsonFormat);
    return responseInJsonFormat

}
const LoadImageAsBase64 = (file) => new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
    reader.readAsDataURL(file);
})

export { GetAllFollower,LoadImageAsBase64 }