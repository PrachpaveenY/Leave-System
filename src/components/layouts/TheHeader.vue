<template>
    <header>
        <h1>{{ title }}</h1>
    </header>
</template>
<script>
export default {
    name: 'TheHeader',
    props: {
        title: String
    }
}
</script>

<style>
header{
    background: #03045E;
}
header h1{
    color: #fff;
    text-align: center;
}
</style>