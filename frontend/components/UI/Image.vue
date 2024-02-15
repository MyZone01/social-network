<!-- ImageUploader.vue -->
<template>
  <div>
    <label>{{ props.label }}</label>
    <input type="file" @change="handleFileChange">
  </div>
</template>

<script setup>

const props = defineProps({
    label: String
})
const selectedImage = ref(null)

const emit = defineEmits(['image-selected'])

const handleFileChange = (event) => {
    const file = event.target.files[0]
    selectedImage.value = file
    emit('image-selected', file)
}
</script>

<style scoped>
input[type="file"] {
    margin-top: 8px;
}
</style>
