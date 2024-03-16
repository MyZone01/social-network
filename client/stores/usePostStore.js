import { defineStore } from 'pinia';

export default defineStore("feed", {
  state: () => ({
    posts: [],
    userPosts: []
  }),
  actions: {
    addPost(post) {
      console.log(post)
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
      console.log("post added: " + comment.post_id);
      for (let i = 0; i < this.posts.length; i++) {
        if (this.posts[i].id === comment.post_id) {
          this.posts[i].comments.push(comment);
          break
        }

      }
    },
    getUserPosts(nickname) {
      let userPosts = []
      for (post of this.posts) {
        if (post.nickname === nickname) {
          userPosts.push(post)
        }
      }
      return userPosts
    }
  }
}); 
