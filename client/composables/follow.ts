export const useFollow = async (nickname: string, action: string) => {
    const response = await fetch(`/api/follow`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': "5d01ea4c-732e-43df-a76f-b493d1df3865",
        },
        body: JSON.stringify({ action: action, nickname: nickname })
    })
    if (response.status !== 200) {
       return false
    }
    return true
}
