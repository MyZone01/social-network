<template>
    <div>
        <label for="" v-if="props.label" class="">
            {{ props.label }}
        </label>
        <div class="mt-2.5">
            <input :value="props.modelValue" @input="event => emits('update:modelValue', event.target.value)" class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" :type="props.type" :placeholder="props.placeholder">
        </div>
    </div>
</template>
<script setup>
const emits = defineEmits(['update:modelValue'])

const props = defineProps({
    modelValue: {
        type: String,
        required: true
    },
    placeholder: {
        type: String,
        default: ''
    },
    type: {
        type: String,
        default: 'text'
    },
    label: {
        type: String,
        default: null
    }
})
</script>