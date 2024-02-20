<template lang="">
    <div class="hidden lg:p-20 uk- open" id="create-status" uk-modal="">

        <div
            class="uk-modal-dialog tt relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[520px] w-full dark:bg-dark2">

            <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
                <h2 class="text-sm font-medium text-black"> Create Post </h2>

                <!-- close button -->
                <button type="button" class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                        stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>

            </div>

            <form action="" ref="post_form">

                <div class="space-y-5 mt-3 p-2">
                    <textarea
                        class="w-full !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl   dark:!text-white dark:placeholder:!text-white dark:!bg-slate-800"
                        name="post-content-text" id="" rows="6" placeholder="What do you have in mind?"></textarea>
                </div>
 
                <div class="flex items-center gap-2 text-sm py-2 px-4 font-medium flex-wrap">
                    <button type="button"
                        class="flex items-center gap-1.5 bg-sky-50 text-sky-600 rounded-full py-1 px-2 border-2 border-sky-100 dark:bg-sky-950 dark:border-sky-900"
                        @click="openFileInput">
                        <ion-icon name="image" class="text-base"></ion-icon>
                        Image
                    </button>
                    <input type="file" id="photo-input" name="photo" accept="image/*" style="display: none"
                        ref="fileInput">


                </div>

                <div class="p-5 flex justify-between items-center">
                    <div>
                        <button
                            class="inline-flex items-center py-1 px-2.5 gap-1 font-medium text-sm rounded-full bg-slate-50 border-2 border-slate-100 group aria-expanded:bg-slate-100 aria-expanded: dark:text-white dark:bg-slate-700 dark:border-slate-600"
                            type="button">
                            Everyone
                            <ion-icon name="chevron-down-outline"
                                class="text-base duration-500 group-aria-expanded:rotate-180"></ion-icon>
                        </button>

                        <div class="p-2 bg-white rounded-lg shadow-lg text-black font-medium border border-slate-100 w-60 dark:bg-slate-700"
                            uk-drop="offset:10;pos: bottom-left; reveal-left;animate-out: true; animation: uk-animation-scale-up uk-transform-origin-bottom-left ; mode:click">

                            <!-- <form>
                        </form> -->
                            <label>
                                <input type="radio" name="radio-status" id="monthly1"
                                    class="peer appearance-none hidden" checked />
                                <div
                                    class=" relative flex items-center justify-between cursor-pointer rounded-md p-2 px-3 hover:bg-secondery peer-checked:[&_.active]:block dark:bg-dark3">
                                    <div class="text-sm"> public </div>
                                    <ion-icon name="checkmark-circle"
                                        class="hidden active absolute -translate-y-1/2 right-2 text-2xl text-blue-600 uk-animation-scale-up"></ion-icon>
                                </div>
                            </label>
                            <label>
                                <input type="radio" name="radio-status" id="monthly1"
                                    class="peer appearance-none hidden" />
                                <div
                                    class=" relative flex items-center justify-between cursor-pointer rounded-md p-2 px-3 hover:bg-secondery peer-checked:[&_.active]:block dark:bg-dark3">
                                    <div class="text-sm"> private </div>
                                    <ion-icon name="checkmark-circle"
                                        class="hidden active absolute -translate-y-1/2 right-2 text-2xl text-blue-600 uk-animation-scale-up"></ion-icon>
                                </div>
                            </label>
                            <label>
                            </label>

                        </div>
                    </div>

                    
                    <UISelecUser />
                    <div class="flex items-center gap-2">
                        <button type="submit" class="button bg-blue-500 text-white py-2 px-12 text-[14px]">
                            Create</button>
                    </div>
                </div>
            </form>

        </div>


    </div>

</template>
<script>
export default {
    mounted() {
        this.$refs.post_form.addEventListener('submit', this.handleSubmitPost)
    },
    methods: {
        openFileInput() {
            this.$refs.fileInput.click();
        },
        handleSubmitPost(e) {
            e.preventDefault();
            let formdata = new FormData(e.target)
            console.log(formdata.get("photo"));

        }
    }
}
</script>
<style lang="">

</style>    