<template>
    <div class="">
        <TheHeader :title="title"/>
        <div class="container">
            <Navigation />
        </div>
    </div>
    
    
</template>

<script>
import TheHeader from './components/layouts/TheHeader.vue';
import Navigation from './components/Auth/Navigation.vue';

export default {
    name : 'App',
    components:{
        Navigation,
        TheHeader,
    },
    data() {
        return {
            title: 'Vue Leave System'
        }
    },
}

</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Nunito+Sans:wght@200&display=swap');

* {
    margin: 0;
    box-sizing: border-box;
    font-family: 'Nunito Sans', sans-serif;
}

/* body, html {
} */
</style>
