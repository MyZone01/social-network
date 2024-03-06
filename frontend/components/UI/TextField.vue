<!-- TextArea.vue -->
<template>
  <div>
    <label>{{ props.label }}</label>
    <textarea :value="props.modelValue" :placeholder="placeholder" rows="4" class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" @input="event => emits('update:modelValue', event.target.value)" />
  </div>
</template>

  <script setup>
  const emits = defineEmits(['update:modelValue'])

  const props = defineProps({
      modelValue: String,
      label: String,
      placeholder: String
    })
  </script>

  <style scoped>
  textarea {
    width: 100%;
    padding: 8px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  </style>
