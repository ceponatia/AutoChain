use std::time::SystemTime;

pub struct Validator {
    pub address: String,
    pub stake: f64,
    pub staking_start: u64,
}

impl Validator {
    pub fn new(address: &str, stake: f64, staking_start: u64) -> Self {
        Validator {
            address: address.to_string(),
            stake,
            staking_start,
        }
    }
}

pub struct WeightedStaking {
    validators: Vec<Validator>,
}

impl WeightedStaking {
    pub fn new(validators: Vec<Validator>) -> Self {
        WeightedStaking { validators }
    }

    pub fn select_validator(&self) -> Option<&Validator> {
        if self.validators.is_empty() {
            return None;
        }

        let current_time = SystemTime::now()
            .duration_since(SystemTime::UNIX_EPOCH)
            .expect("Time went backwards")
            .as_secs();

        let mut selected_validator: Option<&Validator> = None;
        let mut max_weight: f64 = 0.0;

        for validator in &self.validators {
            let held_duration = current_time - validator.staking_start;
            let weight = validator.stake * (held_duration * 2) as f64;

            if weight > max_weight {
                max_weight = weight;
                selected_validator = Some(validator);
            }
        }

        selected_validator
    }
}



