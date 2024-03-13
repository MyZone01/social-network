import { defineStore } from 'pinia';

export const useFeedStore = defineStore("feed", {
    id: 'feed',
    state: () => ({
        posts: []
    }),
    persist: true,
    actions: {
        addPost(post) {
            this.posts.unshift(post);
        },
        async getUserFeed() {
            let response = await fetch("http://localhost:8081/post/getfeed", {
                headers: {
                    Authorization: `Bearer ${useGlobalAuthStore().token}`
                },
                method: 'GET',
            })
                .then((response) => response.json())
                .catch((error) => error)
            this.posts = response
        }
    }
});