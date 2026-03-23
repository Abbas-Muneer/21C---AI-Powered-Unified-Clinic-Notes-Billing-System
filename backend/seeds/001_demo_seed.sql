INSERT INTO patients (id, full_name, date_of_birth, gender, phone, email, address, created_at, updated_at) VALUES
('4f45e8c9-bcdc-4f3e-a78c-6e77d7462a10', 'Nadeesha Fernando', '1991-05-14', 'female', '+94 77 555 1020', 'nadeesha@example.com', '14 Park Road, Colombo', NOW(), NOW()),
('c84410df-2858-4d6c-93a0-6962472b9174', 'Kamal Wijesinghe', '1984-11-30', 'male', '+94 71 555 8844', 'kamal@example.com', '88 Lake View, Kandy', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

INSERT INTO reference_drugs (id, name, default_unit_price, default_route, created_at, updated_at) VALUES
('1dc5f525-f44a-4d25-b7ef-60778b52ca72', 'Amoxicillin', 45, 'oral', NOW(), NOW()),
('707cdbf7-f4eb-4bb0-b153-88f2c9c76902', 'Paracetamol', 15, 'oral', NOW(), NOW()),
('e30111dd-8084-43a7-aeb3-3e76f2f9c6d0', 'Cetirizine', 20, 'oral', NOW(), NOW()),
('8b7fe11d-1e3c-45e7-afc1-c0d8bcf9ccb7', 'Omeprazole', 28, 'oral', NOW(), NOW())
ON CONFLICT (name) DO NOTHING;

INSERT INTO reference_lab_tests (id, name, default_unit_price, created_at, updated_at) VALUES
('1b0db737-4be5-48ad-baa6-a9d90e383e7f', 'Full Blood Count', 1200, NOW(), NOW()),
('0e36d1b0-b99e-4a0b-ac57-471f2dc64875', 'CRP', 1800, NOW(), NOW()),
('7c0e7c75-2b65-4cee-9ded-ff7d6cd00e20', 'Liver Function Test', 2400, NOW(), NOW()),
('58cf8818-13b9-48c4-8b7c-08902bc855fd', 'Urine Full Report', 900, NOW(), NOW())
ON CONFLICT (name) DO NOTHING;
