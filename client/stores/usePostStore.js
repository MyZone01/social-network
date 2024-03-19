import { defineStore } from 'pinia';

export default defineStore("feed", {
  state: () => ({
    posts: [],
    userPosts: []
  }),
  actions: {
    addPost(post) {
      this.posts.unshift(post);
      this.userPosts.unshift(post)
    },
    async getUserFeed() {
      await fetch("/api/getFeed")
        .then(async (response) => {
          const data = await response.json()
          this.posts = data.body
          this.posts.forEach(post => {
            if (post.userOwnerNickname === useAuthUser().value.nickname) {
              this.userPosts.push(post);
            }
          })
        })
        .catch((error) => console.error(error))
    },
    addComment(comment) {
      for (let i = 0; i < this.posts.length; i++) {
        if (this.posts[i].id === comment.post_id) {
          this.posts[i].comments.push(comment);
          break
        }

      }
    },
    getUserPosts(nickname) {
      for (let post of this.posts) {
        console.log(nickname, post)
        if (post.userOwnerNickname === nickname) {
          userPosts.push(post)
        }
      }
      return userPosts
    }
  }
}); 
