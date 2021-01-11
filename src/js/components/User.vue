<template>
  <v-container>
    <v-row>
      <v-col cols="6">
        <h3>Go Echo Sample</h3>
      </v-col>
      <v-col cols="6">
        <p class="text-right mb-0">
          <a @click="logout">Logout</a>
        </p>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <h1>Users</h1>
        <v-btn @click="create">Create</v-btn>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-center users-table-column-id">ID</th>
                <th class="text-center users-table-column-account">Account</th>
                <th class="text-center users-table-column-updated">Updated</th>
                <th class="text-center users-table-column-created">Created</th>
                <th class="text-center users-table-column-buttons"></th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, i) in itemList" :key="i">
                <td class="text-center">{{ item.ID }}</td>
                <td class="font-weight-bold">{{ item.Account }}</td>
                <td class="text-center">{{ item.UpdatedAt | formatDate("YYYY-MM-DD HH:mm") }}</td>
                <td class="text-center">{{ item.CreatedAt | formatDate("YYYY-MM-DD HH:mm") }}</td>
                <td class="text-center">
                  <v-btn @click="openEditDialog(item)">Update</v-btn>
                  <v-btn @click="openDeleteDialog(item)" outlined color="red" class="ml-3">Delete</v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </v-row>
    <v-dialog
      :value="showEditDialog"
      max-width="700"
      transition="dialog-bottom-transition"
      scrollable
      @input="changeShowEditDialog"
    >
      <v-card style="height:max-content">
        <v-card-title>Edit</v-card-title>
        <v-card-text>
          <v-text-field
            type="text"
            v-model="account"
            label="Account"
            :error-messages="errors.Account"
          ></v-text-field>
          <v-text-field
            type="password"
            v-model="password"
            label="Password"
            :error-messages="errors.Password"
          ></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn outlined class="mr-6" @click="closeEditDialog">Cancel</v-btn>
          <v-btn @click="save" v-if="isCreate">Save</v-btn>
          <v-btn @click="update" v-else>Update</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog
      :value="showDeleteDialog"
      max-width="350"
      transition="dialog-bottom-transition"
      scrollable
      persistent
    >
      <v-card style="height:max-content">
        <v-card-title>Delete</v-card-title>
        <v-card-text>Delete OK??</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn outlined class="mr-6" @click="closeDeleteDialog">Cancel</v-btn>
          <v-btn @click="remove" outlined color="red">Delete</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import axios from "axios";
import { mapActions } from "vuex";

export default {
  data() {
    return {
      itemList: [],
      showEditDialog: false,
      id: null,
      account: null,
      password: null,
      isCreate: false,
      errors: {},
      showDeleteDialog: false
    };
  },
  mounted() {
    this.getUsers();
  },
  methods: {
    ...mapActions(["logout", "getAuthHeaders"]),
    async getUsers() {
      const headers = await this.getAuthHeaders();
      axios.get("/api/users", headers)
        .then(response => {
          if (response.status === 200) {
            this.itemList = response.data;
          }
        })
        .catch(error => {
          if (error.response.status === 400 || error.response.status === 401) {
            document.location.href = "/";
          }
          console.log(error.response);
        });
    },
    openEditDialog(item) {
      this.id = item.ID;
      this.account = item.Account;
      this.password = null;
      this.isCreate = false;
      this.errors = {};
      this.showEditDialog = true;
    },
    closeEditDialog() {
      this.showEditDialog = false;
    },
    changeShowEditDialog() {
      this.showEditDialog = !this.showEditDialog;
    },
    create() {
      this.id = null;
      this.account = null;
      this.password = null;
      this.isCreate = true;
      this.errors = {};
      this.showEditDialog = true;
    },
    async save() {
      const url = "/api/users";
      const data = new FormData();
      if (this.account) {
        data.append("account", this.account);
      }
      if (this.password) {
        data.append("password", this.password);
      }

      const headers = await this.getAuthHeaders();
      axios.post(url, data, headers)
        .then(response => {
          if (response.status === 201) {
            console.log("created");
            this.closeEditDialog();
            this.getUsers();
          }
        })
        .catch(error => {
          if (error.response.status === 422) {
            this.errors = error.response.data;
          }
          console.log(error.response);
        });
    },
    async update() {
      const url = "/api/users/" + this.id;
      const data = new FormData();
      if (this.account) {
        data.append("account", this.account);
      }
      if (this.password) {
        data.append("password", this.password);
      }

      const headers = await this.getAuthHeaders();
      axios.patch(url, data, headers)
        .then(response => {
          if (response.status === 200) {
            console.log("updated");
            this.closeEditDialog();
            this.getUsers();
          }
        })
        .catch(error => {
          console.log(error.response);
        });
    },
    openDeleteDialog(item) {
      this.id = item.ID;
      this.showDeleteDialog = true;
    },
    closeDeleteDialog() {
      this.showDeleteDialog = false;
    },
    async remove() {
      const url = "/api/users/" + this.id;
      const headers = await this.getAuthHeaders();
      axios.delete(url, headers)
        .then(response => {
          if (response.status === 200) {
            console.log("deleted");
            this.id = null;
            this.closeDeleteDialog();
            this.getUsers();
          }
        })
        .catch(error => {
          console.log(error.response);
        });
    }
  }
};
</script>

<style scoped>
.users-table-column-id {
  width: 10%;
}
.users-table-column-account {
  width: 40%;
}
.users-table-column-updated {
  width: 20%;
}
.users-table-column-updated {
  width: 20%;
}
.users-table-column-buttons {
  width: 10%;
}
</style>