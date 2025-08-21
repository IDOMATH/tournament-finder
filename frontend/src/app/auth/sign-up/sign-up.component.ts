import { Component } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { valuesMatch } from "../../validators/value-match";

@Component({
  selector: "app-sign-up",
  imports: [],
  templateUrl: "./sign-up.component.html",
  styleUrl: "./sign-up.component.css",
})
export class SignUpComponent {
  form = new FormGroup({
    email: new FormControl({ validators: [Validators.required] }),
    passwords: new FormGroup(
      {
        password: new FormControl({ validators: [Validators.required] }),
        cPassword: new FormControl({ validators: [Validators.required] }),
      },
      { validators: [valuesMatch("password", "cPassword")] }
    ),
  });

  onSubmit() {
    console.log("submitted");
  }
}
