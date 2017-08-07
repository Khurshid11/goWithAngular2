import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from './user';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  newUser: User = new User();

  constructor(private http: HttpClient) { }
  register() {
    console.log("---------------------");
    const body = {
      first_name: this.newUser.firstName,
      last_name: this.newUser.lastName,
      email: this.newUser.email,
      password: this.newUser.password,
    };
    this.http.post('/addUser', body).subscribe(
      data => {

        console.log("data : ", data)
      },
      error => {
        console.log("error : ", error)
      });
  }
}
