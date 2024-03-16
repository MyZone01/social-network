import usePostStore from "~/stores/usePostStore.js";
export const postPreviewContent = ref({
    imageUrl: "",
    comments: []
})
export const passDataOnPostPreviewContent = (postID) => {
    for (let post of usePostStore().posts) {
        if (post.id === postID) {
            postPreviewContent.value = post
            break
        }
    }
}