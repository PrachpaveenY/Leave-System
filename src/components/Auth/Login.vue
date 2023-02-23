<template>
    <page-card>
        <h2>Login Leave System</h2>
        <form action="">
            <div class="form-control">
                <label for="email">Email</label>
                <input type="emal" name="email" id="email">
            </div>
            <div class="form-control">
                <label for="password">Password</label>
                <input type="password" name="password" id="password">
            </div>
            <div>
                <page-button type="submit">
                    Login
                </page-button>
            </div>
        </form>
    </page-card>
</template>
<script>
export default {
    name: 'Login',
}
</script>
<style scoped>
input {
    display: block;
    width: 100%;
    font: inherit;
    padding: 0.15rem;
    border: 1px solid #ccc;
}

input:focus {
    outline: none;
    border-color: #3a0061;
    background-color: #f7ebff;
}

.form-control {
    margin: 1rem 0;
}
</style>