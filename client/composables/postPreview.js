export const postPreviewContent = ref({
    imageUrl: "",
    comments: []
})
export const passDataOnPostPreviewContent = (postID) => {
    for (let post of useFeedStore().posts) {
        if (post.id === postID) {
            postPreviewContent.value = post
            break
        }
    }
}