import { defineStore } from 'pinia';
const { getAllGroupPosts } = useGroupPost()

export default defineStore("feed", {
  state: () => ({
    posts: [],
    userPosts: [],
    groupPosts:[],
    allPosts: new Map()
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
          for (let i = 0; i < this.posts.length; i++) {
            const p = this.posts[i]
            this.allPosts[p.id] = p
            if (this.posts[i].userOwnerNickname === useAuthUser().value.nickname) {
              this.userPosts.push(this.posts[i]);
            }
          }
        })
        .catch((error) => console.error(error))
    },

    async getGroupFeeds(groupId) {
      this.groupPosts = await getAllGroupPosts(groupId)
      this.groupPosts.forEach(p=>{
        this.allPosts[p.id] = p
      })
    },

    addComment(comment) {
      console.log(comment);
      // for (let i = 0; i < this.posts.length; i++) {
      //   if (this.posts[i].id === comment.post_id) {
      //     this.posts[i].comments.push(comment);
      //     break
      //   }

      // }

      this.allPosts[comment.post_id].comments.push(comment)

    },
    getpostComments(postId){
      return this.allPosts[postId].comments
    },
    getUserPosts(nickname) {

      let userPosts = []
      for (let post of this.posts) {
        console.log(nickname, post)
        if (post.userOwnerNickname === nickname) {
          userPosts.push(post)
        }
      }
      return userPosts
    },
    flushAllPosts() {
      this.userPosts = []
      this.posts = []
      this.groupPosts = []
      this.allPosts = []
    }
  }
}); 
