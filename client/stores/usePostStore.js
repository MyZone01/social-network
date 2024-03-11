import { defineStore } from 'pinia';

export const useFeedStore = defineStore("feed", {
    id: 'feed',
    state: () => ({
        posts: []
    }),
    persist: true,
    actions: {
        addPost(post) {
            console.log(post)
            this.posts.unshift(post);
        },

        async getUserFeed() {
            await fetch("/api/getFeed")
                .then(async (response) => {
                    const data  = await response.json()
                    this.posts= data.body
                })
                .catch((error) => console.error(error))
            console.log(this.posts);
        }
    }
}); 