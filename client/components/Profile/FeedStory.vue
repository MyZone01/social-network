import Detail from './Detail.vue';

<template>
    <div class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col" id="js-oversized">

        <!-- feed story -->

        <!-- posts component  -->

        <Detail v-if="status == 'about'" :data="data" />

    </div>
</template>

<script>
import Detail from './Detail.vue';
export default {
    name: 'FeedStory',
    components: {
        Detail
    },
    props: {
        status: {
            type: String,
            required: true
        },
        data: {
            type: Object,
            required: true
        },
        // posts: {
        //     type: Array,
        //     required: true
        // }
    },
}
</script>