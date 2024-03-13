const GetAllFollower = async () => {
    let response = await fetch("/api/getFollowers")
    if (response.statusCode != 200 ){
        return {status: response.statusCode}
    }
    let responseInJsonFormat = await response.json().catch(err => ({ error: err }))
    return responseInJsonFormat

}


export { GetAllFollower }