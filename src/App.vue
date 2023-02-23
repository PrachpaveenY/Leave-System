<template>
    <div class="">
        <TheHeader :title="title"/>
        <div class="container">
         <page-card>
           <page-button @click="switchTab('login')" :mode="setLoginActiveTab">Login</page-button>
           <page-button @click="switchTab('register')" :mode="setRegisterActiveTab">Register</page-button>
         </page-card>
         <keep-alive>
            <component :is="selectedTab"></component>
         </keep-alive>
        </div>
    </div>
    
    
</template>

<script>
import TheHeader from './components/layouts/TheHeader.vue';
import Login from './components/Auth/Login.vue';
import Register from './components/Auth/Register.vue';

export default {
    name : 'App',
    components:{
        TheHeader,
        Login,
        Register,
    },
    data() {
        return {
            title: 'Vue.JS Leave System',
            selectedTab : Login,
        }
    },
    methods:{
        switchTab(tab){
            this.selectedTab = tab;
        }
    },
    computed:{
        // เซ็ตเป็นค่าว่างเมื่อคลิก
        setLoginActiveTab(){
            return this.selectedTab !== 'login' ? '' : 'default';
        },
        setRegisterActiveTab(){
            return this.selectedTab !== 'register' ? '' : 'default';
        }
    }
}

</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Nunito+Sans:wght@200&display=swap');

* {
    margin: 0;
    box-sizing: border-box;
    font-family: 'Nunito Sans', sans-serif;
}

.container {
    width: 90%;
    margin: 0 auto;
}

/* body, html {} */
</style>
