export const useFollow = async (nickname: string, action: string) => {
    const response = await fetch(`/api/follow`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ action: action, nickname: nickname })
    })
    if (response.status !== 200) {
       return false
    }
    return true
}
