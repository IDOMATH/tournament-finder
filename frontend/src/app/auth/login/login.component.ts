import { Component } from "@angular/core";
import {
  AbstractControl,
  FormControl,
  FormGroup,
  Validators,
} from "@angular/forms";

function valuesMatch(controlName1: string, controlName2: string) {
  return (control: AbstractControl) => {
    const val1 = control.get(controlName1)?.value;
    const val2 = control.get(controlName2)?.value;

    if (val1 === val2) {
      return null;
    }

    return { passwordsDoNotMatch: true };
  };
}

@Component({
  selector: "app-login",
  imports: [],
  templateUrl: "./login.component.html",
  styleUrl: "./login.component.css",
})
export class LoginComponent {
  form = new FormGroup({
    email: new FormControl({ validators: [Validators.required] }),
    password: new FormControl({ validators: [Validators.required] }),
  });

  onSubmit() {
    console.log("Submitted");
  }
}
