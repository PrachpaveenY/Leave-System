<template lang="">
    <button :mode="mode">
        <slot />
    </button>
</template>
<script>
export default {
    name: 'PageButton',
    props: {
        mode: String,
    }
}
</script>
<style scoped>
    button {
        padding: 0.75rem 1.5rem;
        font-family: inherit;
        background-color: #03045E;
        border: 1px solid #03045E;
        color: white;
        cursor: pointer;
    }

    button:hover,
    button:active {
        background-color: #023E8A;
        border-color: #023E8A;
    }

    .default {
        background-color: #03045E;
        color: #03045E;
        border: none;
    }

    .default:hover,
    .default:active {
        background-color: #eee;
    }
</style>