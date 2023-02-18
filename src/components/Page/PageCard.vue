<template>
    <div>
        <slot />
    </div>
</template>
<script>
export default {
    name: 'PageCard',
}
</script>
<style scoped>
    div {
        border-radius: 12px;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
        padding: 1 rem;
        margin: 2rem auto;
        max-width: 40rem;
    }
</style>