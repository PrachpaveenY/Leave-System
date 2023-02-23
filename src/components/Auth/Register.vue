<template>
    <page-card>
        <h2>Register Leave System</h2>
        <form action="">
            <div class="form-control">
                <label for="fullname">Fullname</label>
                <input type="text" name="fullname" id="fullname">
            </div>
            <div class="form-control">
                <label for="email">Email</label>
                <input type="emal" name="email" id="email">
            </div>
            <div class="form-control">
                <label for="password">Password</label>
                <input type="password" name="password" id="password">
            </div>
            <div class="form-control">
                <label for="password2">Confirm-Password</label>
                <input type="password" name="password2" id="password2">
            </div>
            <div>
                <page-button type="submit">
                    Register
                </page-button>
            </div>
        </form>
    </page-card>
</template>
<script>
export default {
    name: 'Register',
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