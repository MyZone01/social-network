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
          for (let i = 0; i < this.posts.length; i++){
            if (this.posts[i].userOwnerNickname === useAuthUser().value.nickname) {
              this.userPosts.push(this.posts[i]);
            }
          }
        })
        .catch((error) => console.error(error))
    },
    addComment(comment) {
      console.log(comment);
      for (let i = 0; i < this.posts.length; i++) {
        if (this.posts[i].id === comment.post_id) {
          
          this.posts[i].comments.push(comment);
          break
        }

      }
    },
    getUserPosts(nickname) {

      let userPosts = []
      for (let post of this.posts) {
        
        if (post.userOwnerNickname === nickname) {
          userPosts.push(post)
        }
      }
      return userPosts
    }
  }
}); 
